from setuptools import setup, find_packages

setup(
    name='tttime',
    version='1.0.0',
    author='tonyaellie',
    author_email='tonya@tokia.dev',
    description='An implementation of a TTTime in Python.',
    long_description=open('../README.md').read(),
    long_description_content_type='text/markdown',
    url='https://github.com/tonyaellie/tttime',
    packages=find_packages(),
    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)