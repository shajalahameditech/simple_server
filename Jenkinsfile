pipeline {
    agent {label 'jenkins_slave'}

    environment {
          DOCKER_HUB_ACCOUNT="shajal"
          DOCKER_IMAGE_NAME="webpage"
          K8S_DEPLOYMENT_NAME="webapp"
       }

    stages {
        stage('Docker Build and push') {
            steps {

                echo 'Building..'
                sh ("sudo docker build -t ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} .")
                sh ("sudo docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER}")
                sh ("sudo docker tag ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")
                sh ("sudo docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")

            }
        }
    }

}