import { TimeEntry, TimeEntryRepository } from '@/businesslogic';
import { projectConfig } from '@/config'
import firebase from 'firebase/app'

export class TimeEntryRepositoryFirestore implements TimeEntryRepository {
  async addTimeEntry(te: TimeEntry): Promise<string> {
    const colRef = this.getTimeEntriesColRef()
    const result = await colRef.add(te)
    return result.id
  }

  async getAllTimeEntries(): Promise<TimeEntry[]> {
    const snapshot = await this.getTimeEntriesColRef().orderBy('startTime', 'desc')
      .get()
    if (snapshot.empty) return []
    return snapshot.docs.map(doc => doc.data())
  }

  private getTimeEntriesColRef() {
    const uid = firebase.auth().currentUser?.uid
    if (!uid) throw new Error('Cannot get current user UID')
    const docPath = projectConfig.timetracking.firestoreTimeEntryCol(uid)
    return firebase
      .firestore()
      .collection(docPath)
      .withConverter(timeEntryConverter)
  }

  private getTimeEntryDocRef(teId: string) {
    return this.getTimeEntriesColRef()
      .doc(teId)
      .withConverter(timeEntryConverter)
  }
}


const timeEntryConverter = {
  toFirestore(te: TimeEntry): firebase.firestore.DocumentData {
    return te;
  },
  fromFirestore(
    snapshot: firebase.firestore.QueryDocumentSnapshot,
    options: firebase.firestore.SnapshotOptions
  ): TimeEntry {
    const data: any = snapshot.data(options)!;
    return {
      ...data,
      id: snapshot.id
    }
  }
}
