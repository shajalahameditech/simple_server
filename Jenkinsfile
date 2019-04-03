pipeline {
    agent {label 'jenkins_slave'}

    stages {
        stage('Build') {
            steps {
                echo 'Building..'

            }
        }
    }

    post {
        success {
            echo 'This will run only if successful'
        }
    }
}