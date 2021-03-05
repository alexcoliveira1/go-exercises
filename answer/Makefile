build:
	docker build -t answer-container -f Dockerfile .

run:
	docker run -d -p 8090:8080 --rm --name=answerapi answer-container

stop:
	docker stop answerapi
