pipeline {
  environment {
    registry = "shajalahamedcse/webpage"
    registryCredential = 'dockerhub'
    dockerImage = ''
  }
  agent {label 'kubernetes-deploy'}
  stages {
    stage('Cloning Git') {
      steps {
        git 'https://github.com/shajalahameditech/simple_server'
      }
    }
    stage('Building image') {
      steps{
        script {
          dockerImage = docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }
    stage('Deploy Image') {
      steps{
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push()
          }
        }
      }
    }
    stage('Remove Unused docker image') {
      steps{
        sh "docker rmi $registry:$BUILD_NUMBER"
      }
    }
  }
}