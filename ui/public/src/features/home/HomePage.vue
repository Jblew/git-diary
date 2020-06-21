<template>
  <div>
    <h1>Ephemeris, {{ user.displayName }}</h1>
    <router-link to="/logout">Log out</router-link>
    <hr />

    <diary-panel />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { DiaryPanel } from '@/features/diary';

@Component({
  components: {
    DiaryPanel,
  },
})
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;

  /*// tslint:disable no-console
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
  }*/
}
</script>

<style scoped>
</style>
