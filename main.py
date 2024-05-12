import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'

import io
import tensorflow as tf
from tensorflow import keras
import numpy as np
from PIL import Image
import sys

model = keras.models.load_model("nn.h5")

def transform_image(pillow_image):
    data = np.asarray(pillow_image)
    data = data / 255.0
    data = data[np.newaxis, ..., np.newaxis]
    # --> [1, x, y, 1]
    data = tf.image.resize(data, [28, 28])
    return data


def predict(x):
    predictions = model(x)
    predictions = tf.nn.softmax(predictions)
    pred0 = predictions[0]
    label0 = np.argmax(pred0)
    return label0

def main():
    filepath = sys.argv[1]
    # filepath = "./image/tujuh.png"
    with open(filepath, "rb") as file:
        image_bytes = file.read()
    pillow_img = Image.open(io.BytesIO(image_bytes)).convert('L')
    tensor = transform_image(pillow_img)
    prediction = predict(tensor)
    data = {"prediction": int(prediction)}
    print(data)


if __name__ == "__main__":
    main()