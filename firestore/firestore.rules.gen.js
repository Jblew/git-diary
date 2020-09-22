const fs = require('fs')

function getRules() {
  return `
rules_version = '2';

service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow write: if false;
      allow read: if false;
    }

    ///////////////////////
    // Utility functions //
    ///////////////////////

    function isAuthenticated() {
        return request.auth != null;
    }

    ${rulesForEnv('test')}
    ${rulesForEnv('prod')}
  }
}
  `
}

function rulesForEnv(envName) {
  const basePath = `/databases/$(database)/documents/envs/${envName}`
  return `
  match /envs/${envName} {
    ${timeTracking()}
  }
  `
}

function timeTracking() {
  return `
    //////////////////////////
    // Time tracking system //
    //////////////////////////

    match /timetracking/{uid}/entries {
      allow write: if isAuthenticated() && uid == request.auth.uid;
      allow read: if isAuthenticated() && uid == request.auth.uid;
    }
  `
}

const data = getRules()
fs.writeFileSync(__dirname + '/firestore.rules', data)
console.log('Regenerated firestore rules')
