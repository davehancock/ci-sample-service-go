pipeline {
    agent none

    environment {
        IMAGE = 'daves125125/ci-sample-service-go'
        DOCKER_LOGIN = credentials('docker-registry-login')
    }

    stages {

        stage('Build') {
            agent { docker 'golang:1.9' }
            steps {
                sh 'mkdir -p ${GOPATH}/src/github && ln -s ${GOPATH}/src/github/ci-sample-service-go ${WORKSPACE}'

                sh 'go get -u github.com/golang/dep/cmd/dep'
                sh 'cd ${GOPATH}/src/github/ci-sample-service-go && ls -ltra && dep ensure'

                sh 'cd ${GOPATH}/src/github/ci-sample-service-go && env GOOS=linux GOARCH=386 go build -o ci-sample-service-go'
            }
        }

        stage('Test') {
            agent { docker 'golang:1.9' }
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Deploy Snapshot') {
            agent any
            steps {
                script {
                    def HASH = sh returnStdout: true, script: 'git rev-parse HEAD'
                    sh """
                        docker login -u ${DOCKER_LOGIN_USR} -p ${DOCKER_LOGIN_PSW}
                        docker build -t ${IMAGE} .
                        docker tag ${IMAGE} ${IMAGE}:${HASH}
                        docker push ${IMAGE}:${HASH}
                    """
                }
                // FIXME put in post section
                deleteDir()
            }
        }

        stage('Deploy Release') {
            agent any
            when {
                branch 'master'
            }
            steps {
                sh """
                    docker tag ${IMAGE} ${IMAGE}:latest
                    docker push ${IMAGE}:latest
                """
            }
        }
    }

}
