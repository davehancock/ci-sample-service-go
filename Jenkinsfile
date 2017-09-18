pipeline {
    agent none

    environment {
        IMAGE = 'daves125125/ci-sample-service-go'
        DOCKER_LOGIN = credentials('docker-registry-login')
    }

    stages {

        stage('Build') {
            agent { docker 'golang:1.9' }
            environment {
                GOPATH = "${WORKSPACE}"
            }
            steps {
                sh """
                    wget https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-386.zip
                    gunzip -S .zip -c dep-linux-386.zip > ./dep && chmod 755 dep
                """
                sh './dep ensure'
                sh 'env GOOS=linux GOARCH=386 go build -o ci-sample-service-go'
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
