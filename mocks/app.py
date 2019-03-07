#  -*- encoding: utf-8 -*-
import argparse
import logging
import traceback
import time

from flask import Flask, jsonify

app = Flask(__name__)


logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                    level=logging.os.environ.get('LOGGING_LEVEL'))
logger = logging.getLogger(__name__)


def get_arg(args=None):
    parser = argparse.ArgumentParser()
    parser.add_argument("-host", "--arg_host", required=True)
    parser.add_argument("-port", "--arg_port", type=int, required=True)
    results = parser.parse_args(args)
    return results.arg_host, results.arg_port


host, port = get_arg()


@app.route("/fast-ok")
def fast():
    try:
        return jsonify('OK'), 200
    except Exception:
        logging.error("Something went wrong. Exception: {exception}".format(exception=traceback.format_exc()))


@app.route("/fast-error")
def fast_error():
    try:
        return jsonify('Internal Server Error'), 500
    except Exception:
        logging.error("Something went wrong. Exception: {exception}".format(exception=traceback.format_exc()))


@app.route("/slow-ok/<sleep_time>")
def slow(sleep_time):
    try:
        time.sleep(int(sleep_time))
        return jsonify('OK'), 200
    except Exception:
        logging.error("Something went wrong. Exception: {exception}".format(exception=traceback.format_exc()))


@app.route("/slow-error/<sleep_time>")
def slow_error(sleep_time):
    try:
        time.sleep(int(sleep_time))
        return jsonify('Internal Server Error'), 500
    except Exception:
        logging.error("Something went wrong. Exception: {exception}".format(exception=traceback.format_exc()))


if __name__ == '__main__':
    app.run(host=host, port=port, debug=False)
