<template>
  <main>
    <h1> Teams </h1>
  </main>
</template>

<script>
import config from '../config';

export default {
  name: 'Teams',
  mounted: () => {
    fetch(`${config.SERVER_URL}/api/teams`)
      .then(response => response.json())
      .then(console.log);
  },
};
</script>

<style>
</style>
