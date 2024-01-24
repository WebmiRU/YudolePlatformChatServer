<script setup>
</script>

<template>
  <header>

  </header>

  <div class="columns">
  <aside class="menu column is-4">
    <p class="menu-label">
      Чаты
    </p>
    <ul class="menu-list">
<!--      <a class="is-active">GoodGame</a>-->
      <li><router-link :to="{name: 'chats.goodgame'}" active-class="is-active">GoodGame</router-link></li>
      <li><router-link :to="{name: 'chats.trovo'}" active-class="is-active">Trovo</router-link></li>
      <li><router-link :to="{name: 'chats.twitch'}" active-class="is-active">Twitch</router-link></li>
      <li><router-link :to="{name: 'chats.youtube'}" active-class="is-active">YouTube</router-link></li>
    </ul>
    <p class="menu-label">
      Прочие параметры
    </p>
    <ul class="menu-list">
      <li><a>Team Settings</a></li>
      <li><a>Manage Your Team</a></li>
    </ul>

  </aside>
  <main class="column is-8">
    <router-view></router-view>
  </main>
  </div>
</template>
