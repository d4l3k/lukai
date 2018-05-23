from setuptools import setup, find_packages
import os

setup(
    name='lukai',
    packages=find_packages('.'),
    version='0.4',
    description='Luk.ai management library',
    install_requires=['grpcio', 'six', 'srvlookup'],
    author='Tristan Rice',
    author_email='rice@fn.lc',
    url='https://github.com/luk-ai/lukai/tree/master/py',
    keywords=['ml','machine learning','lukai'],
    classifiers=[],
)
