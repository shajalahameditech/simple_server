pipeline {
    agent {label 'jenkins_slave'}

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh " kubectl get nodes"
            }
        }
    }

    post {
        success {
            echo 'This will run only if successful'
        }
    }
}