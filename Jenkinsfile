pipeline {
    agent any

    environment {
        IMAGE = "alidevops8/hello-devops:latest"
    }

    stages {
        
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                sh """
                docker build -t $IMAGE ./app
                """
            }
        }
        
        stage('Push Docker Image') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh """
                    echo $\DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    docker push $IMAGE
                    """
                }
            }
        }

        stage('Notify ArgoCD') {
            steps {
                echo "Trigger ArgoCD will be added later"
            }
        }
    }

    post {
        success {
            echo "🚀 Build complete and pushed to Docker Hub!"
        }
        failure {
            echo "❌ Pipeline failed"
        }
    }
}
