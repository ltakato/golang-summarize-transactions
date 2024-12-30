import logging
import pandas as pd
import torch
from transformers import pipeline
from flask import Flask, request, jsonify
import base64
from google.cloud import storage

app = Flask(__name__)

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger()

classifier = pipeline("zero-shot-classification", model="facebook/bart-large-mnli")

CATEGORIES = ["food", "transport","shopping", "entertainment", "grocery"]

@app.route('/user-categories-pubsub', methods=['POST'])
def pubsub_trigger():
    envelope = request.get_json()
    if not envelope:
        return 'Bad Request: No Pub/Sub message received', 400

    process()

    pubsub_message = envelope.get('message', {})
    data = pubsub_message.get('data', '')
    if data:
        data = base64.b64decode(data).decode('utf-8')
        return f"Received message: {data}", 200
    else:
        return 'Bad Request: No message data found', 400

def process():
    try:
        df = pd.read_csv('extract-example.csv')
        classifications = []
        for title in df['title']:
            result = classifier(title, CATEGORIES)
            predicted_label = result['labels'][0]
            score = result['scores'][0]

            logger.info(f"Classified title: '{title}' as '{predicted_label}' with score: {score:.4f}")

            classifications.append({
                "title": title,
                "category": predicted_label,
                "score": score
            })
    except Exception as e:
        return jsonify({"error": f"Failed to process CSV file: {str(e)}"}), 400

def download_storage_file(bucket_name, blob_name, destination_file_name):
    storage_client = storage.Client()

    bucket = storage_client.get_bucket(bucket_name)

    blob = bucket.blob(blob_name)

    blob.download_to_filename(destination_file_name)

    print(f"File {blob_name} downloaded to {destination_file_name}.")

def read_and_print_csv(local_file_path):
    df = pd.read_csv(local_file_path)
    print(df)

if __name__ == '__main__':
    bucket_name = 'summary-transactions'
    blob_name = 'extract-example.csv'
    destination_file_name = '/tmp/extract-example.csv'  # Local file path to save

    download_storage_file(bucket_name, blob_name, destination_file_name)
    print('downloaded from storage')
    read_and_print_csv(destination_file_name)

    app.run(debug=True, host='0.0.0.0', port=8080)
