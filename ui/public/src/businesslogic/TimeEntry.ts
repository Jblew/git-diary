export interface TimeEntry {
  id?: string
  name: string
  startTime: Date
  endTime?: Date
  tags: string[]
  project: string
  comment: string
}
