import { isE2ETest } from '@/util'
import firebase from 'firebase/app'

const env = isE2ETest() ? 'test' : 'prod'
const envCol = `envs/${env}`

export const projectConfig = {
  firebaseAuth: {
    signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],
  },
  timetracking: {
    firestoreTimeEntryCol: (uid: string) => `${envCol}/timetracking/${uid}/entries`,
  }
};
