# =============================================================================
# Make file for development and testing poole
#
# You most likely will run poole.py in your own virtual env.
# =============================================================================

export LC_ALL := en_US.UTF-8
export PYTHONIOENCODING := UTF-8:replace

HERE:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := build

# =============================================================================
# clean up
# =============================================================================

.PHONY: clean
clean:
	rm -f *.pyc
	touch requirements.txt

.PHONY: distclean
distclean: clean
	rm -rf env

# =============================================================================
# build virtual environment
# =============================================================================

env/bin/python:
	python3 -m venv env

env/_requirements: requirements.txt
	env/bin/pip install --upgrade pip
	env/bin/pip install -U -r requirements.txt
	touch $@

.PHONY: build
build: ## build everything needed to run tests and deploy releases
build: env/bin/python env/_requirements


# =============================================================================
# tests
# =============================================================================

.PHONY: test
test: ## run the test suite
test: build
	cd tests && $(HERE)/env/bin/python run.py

