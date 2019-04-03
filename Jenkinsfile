pipeline {
    agent {label 'jenkins_slave'}
    def DOCKER_HUB_ACCOUNT = 'smesch'
    def DOCKER_IMAGE_NAME = 'hugo-app-jenkins'
    def K8S_DEPLOYMENT_NAME = 'hugo-app'
    environment {
          DOCKER_HUB_ACCOUNT="shajal"
          DOCKER_IMAGE_NAME="webpage"
          K8S_DEPLOYMENT_NAME="webapp"
       }

    stages {
        stage('Docker Build and push') {
            steps {
                echo 'Building..'
                sh ("docker build -t ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} .")
                sh ("docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER}")
                sh ("docker tag ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")
                sh ("docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")

            }
        }
    }

}