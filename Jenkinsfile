/* Requires the Docker Pipeline plugin */
pipeline {
    agent {
        docker {
            image 'golang:1.18'
            label 'master'
        }
    }

    parameters {
//         string description: 'specify credentials id', name: 'CREDS', trim: true
        credentials credentialType: 'com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey', defaultValue: 'gp0501', name: 'CREDS', required: true
        string defaultValue: 'main', description: 'specify the branch to checkout', name: 'BRANCH', trim: true
        booleanParam defaultValue: true, description: 'specify the result of test', name: 'TEST_PASS'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: "${BRANCH}"]],
                    userRemoteConfigs: [[url: "$GIT_URL"]]
                ])
            }
        }
        stage('Build') {
            steps {
                sh 'git branch'
                sh 'CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o myjenkins main.go'
                archiveArtifacts artifacts: 'myjenkins', fingerprint: true
            }
        }
        stage('Test') {
            steps {
                sh 'echo "TEST_PASS is ${TEST_PASS}"'
                sh './myjenkins --test=${TEST_PASS}'
//                 junit '**/target/*.xml'
            }
        }
        stage('Deploy') {
            environment {
                SSH_CREDS = credentials("${CREDS}")
            }
            when {
              expression {
                currentBuild.result == null || currentBuild.result == 'SUCCESS'
              }
            }
            steps {
                sh 'echo "SSH private key is located at $SSH_CREDS"'
                sh 'echo "SSH user is $SSH_CREDS_USR"'
                sh 'echo "SSH passphrase is $SSH_CREDS_PSW"'
                sh 'scp -o StrictHostKeyChecking=no -i $SSH_CREDS myjenkins $SSH_CREDS_USR@172.16.10.58:~/'
//              sshagent(["${CREDS}"]) {
//                  sh "scp -o StrictHostKeyChecking=no myjenkins starunion@172.16.10.58:~/"
//              }
            }
        }
    }

    post {
//         always {
//             junit '**/target/*.xml'
//         }

        failure {
            mail to: '188611208@qq.com', subject: 'The Pipeline failed :(', body:'123'
        }
    }
}