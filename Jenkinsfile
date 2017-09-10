pipeline {
    agent none

    environment {
        IMAGE = 'daves125125/ci-sample-service-go'
    }

    stages {

        stage('Build') {
            agent { docker 'golang:1.9' }
            steps {
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
