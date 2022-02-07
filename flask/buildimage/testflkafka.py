import logging
from joblib import load
from flask import Flask, request
import pandas as pd

from io import StringIO
import sys

from kafka import KafkaProducer

bootstrap_servers = ['kafka-1.default.svc.cluster.local:9092']
topicName='test1'
producer = KafkaProducer(bootstrap_servers = bootstrap_servers)


INPUT_ARRAY = [[5.1, 3.5, 1.4, 0.2]]
logging.basicConfig(level=logging.DEBUG)

app = Flask(__name__)
clf = load('/home/svc_model.model')

@app.route('/')
def hello_world():
   clf = load('/home/svc_model.model')
   preds = clf.predict(INPUT_ARRAY)
   app.logger.info(" Inputs: {}".format(INPUT_ARRAY))
   app.logger.info(" Prediction: {}".format(preds))
   message = producer.send(topicName, str(preds).encode('utf-8'))
   tosend = message.get(timeout=10)
   producer.flush()
   #return str(preds)
   return {"prediction": str(preds)}

@app.route('/predict', methods=['POST'])
def predict():

    data = request.get_json()
    app.logger.info("Record To predict: {}".format(data))
    app.logger.info(type(data))
    input_data = [data["data"]]
    app.logger.info(input_data)
    prediction = clf.predict(input_data)
    app.logger.info(prediction)
    response_data = prediction[0]

    message_to_kafka = producer.send(topicName, str(response_data).encode('utf-8'))
    tosend = message_to_kafka.get(timeout=20)

    producer.flush()

    return {"prediction": str(response_data)}

@app.route('/kafka', methods=['GET'])
def sending():

    message = producer.send(topicName, b'henlo from flask')
    tosend = message.get(timeout=10)
    producer.flush()

    return "sent"


if __name__ == '__main__':
    app.run(host='0.0.0.0',  debug=True)
