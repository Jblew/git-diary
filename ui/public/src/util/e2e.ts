import { getCookie } from './cookie';

export function isE2ETest() {
  return getCookie("e2e") === "1"
}
