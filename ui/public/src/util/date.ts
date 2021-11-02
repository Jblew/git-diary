export function dateToISODate(date: Date) {
  const monthStartWithOne = date.getMonth() + 1
  const y = date.getFullYear()
  const m = `${monthStartWithOne}`.padStart(2, '0')
  const d = `${date.getDate()}`.padStart(2, '0')
  return `${y}-${m}-${d}`
}

export function dateToISOTime(date: Date, { showSeconds }: { showSeconds?: boolean } = {}) {
  const h = `${date.getHours()}`.padStart(2, '0')
  const m = `${date.getMinutes()}`.padStart(2, '0')
  const s = `${date.getSeconds()}`.padStart(2, '0')
  if (showSeconds) {
    return `${h}:${m}:${s}`

  }
  return `${h}:${m}`
}
