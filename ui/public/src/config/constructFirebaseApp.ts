import * as firebase from 'firebase/app';
import 'firebase/firestore';
import 'firebase/auth';
import 'firebase/database';

const FIREBASE_CONFIG = {
  apiKey: 'AIzaSyBIWdkm7VQ2YleeJ8YC-ux7JB560JyeeBc',
  authDomain: 'git-diary.firebaseapp.com',
  databaseURL: 'https://git-diary.firebaseio.com',
  projectId: 'git-diary',
  storageBucket: 'git-diary.appspot.com',
  messagingSenderId: '277824411111',
  appId: '1:277824411111:web:b8c6101cd84fa830535942'
};

export function constructFirebaseApp() {
  return firebase.initializeApp(FIREBASE_CONFIG);
}
