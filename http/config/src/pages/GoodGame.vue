<script setup>
import {config} from "../store/config.js";
</script>

<template>
  <h1 class="title is-1">GoodGame</h1>
  <div class="field">
    <label class="label">Имя пользователя</label>
    <div class="control">
      <input v-model="config.services.goodgame.login" class="input" type="text" placeholder="Имя пользователя">
    </div>
  </div>

  <div class="field">
    <label class="label">Пароль</label>
    <div class="control">
      <input v-model="config.services.goodgame.password" class="input" type="text" placeholder="Пароль">
    </div>
  </div>

  <div class="field">
    <label class="label">Каналы</label>
  </div>

  <div v-for="(v, i) in config.services.goodgame.channels" class="field has-addons">
    <div class="control is-expanded">
      <input v-model="config.services.goodgame.channels[i]" class="input" type="number" min="1" max="999999999"
             placeholder="Номер канала">
    </div>
    <div class="control" v-if="config.services.goodgame.channels.length > 1">
      <a @click="config.services.goodgame.channels.splice(i, 1)" class="button is-danger">-</a>
    </div>
    <div class="control">
      <a @click="config.services.goodgame.channels.push(null)" class="button is-primary">+</a>
    </div>
  </div>
</template>
