<template>
  <div>
    <h1>Welcome {{ user.displayName }}</h1>
    <router-link to="/logout">Log out</router-link>
    <hr />
    <pre>
      {{ fetchResponse }}
    </pre>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { functions } from 'firebase';

@Component({
  components: {
  },
})
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public fetchResponse: string = '';

  public user: firebase.User = this.firebase.auth().currentUser!;

  // tslint:disable no-console
  public mounted() {
    this.fetchResponse = 'Loading...';

    // this.firebase.functions().httpsCallable('PublishEntry')()
    this.publishEntry()
      .then(
        (resp) => {
          console.log(resp);
          this.fetchResponse = 'Done.';
        },
        (err) => {
          console.error(err);
          this.fetchResponse = 'Error: ' + err.message;
        },
      );
  }

  private async publishEntry(): Promise<string> {
    const token = await this.user.getIdToken();
    const response = await fetch('/publish', {
      headers: {
        Authorization: `bearer ${token}`,
      },
    });
    return response.text();
  }
}
</script>

<style scoped>
</style>
