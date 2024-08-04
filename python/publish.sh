#!/bin/bash

python publish.py
python setup.py sdist bdist_wheel
twine upload dist/*