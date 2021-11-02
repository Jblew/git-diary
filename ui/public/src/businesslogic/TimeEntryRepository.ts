import { TimeEntry } from './TimeEntry';

export interface TimeEntryRepository {
  addTimeEntry(e: TimeEntry): Promise<string>
  getAllTimeEntries(): Promise<TimeEntry[]>
}
