<script setup>
import {config} from "../store/config.js";
</script>

<template>
  <h1 class="title is-1">Twitch</h1>

  <div class="field">
    <label class="label">Имя пользователя</label>
    <div class="control">
      <input v-model="config.services.twitch.login" class="input" type="text" placeholder="Имя пользователя">
    </div>
  </div>

  <div class="field">
    <label class="label">Пароль</label>
    <div class="control">
      <input v-model="config.services.twitch.password" class="input" type="text" placeholder="Пароль">
    </div>
  </div>

  <div class="field">
    <label class="label">Каналы</label>
  </div>

  <div class="field">
    <a @click="config.services.twitch.channels.push(null)" class="button is-primary">Добавить канал</a>
  </div>

  <div v-for="(v, i) in config.services.twitch.channels" class="field has-addons">
    <div class="control is-expanded">
      <input v-model="config.services.twitch.channels[i]" class="input" type="text" placeholder="Канал">
    </div>
    <div class="control">
      <a @click="config.services.twitch.channels.splice(i, 1)" class="button is-danger">-</a>
    </div>
  </div>
</template>
