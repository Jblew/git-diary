import firebase from 'firebase/app';

export async function readDiary(): Promise<string> {
  const response = await fetch('/read', {
    headers: await getAuthHeaders(),
  });
  const responseText = await response.text();
  validateErrors(responseText);
  return responseText;
}

export async function publishEntry(entry: string): Promise<string> {
  const response = await fetch('/publish', {
    method: 'POST',
    headers: await getAuthHeaders(),
    body: entry,
  });
  const responseText = await response.text();
  validateErrors(responseText);
  return responseText;
}

async function getAuthHeaders() {
  const user = firebase.auth().currentUser;
  if (!user) {
    throw new Error('Cannot retrive current user');
  }
  const token = await user.getIdToken();

  return {
    Authorization: `bearer ${token}`,
  };
}

async function validateErrors(response: string) {
  let jsonObj: any = {};
  try {
    jsonObj = JSON.parse(response);
  } catch (err) {
    // response is not json
  }
  if (jsonObj.error) {
    throw new Error(jsonObj.error);
  }
}
