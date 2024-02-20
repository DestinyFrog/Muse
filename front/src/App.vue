<script>
import TheHeader from './components/TheHeader.vue'
import SongList from './components/song/SongList.vue'
import TheFooter from './components/TheFooter.vue'

export default {
	name: 'App',
	components: {
		TheHeader,
		SongList,
		TheFooter,
	},
}
</script>

<template>
	<TheHeader/>
	<SongList/>
	<TheFooter :SelectedSong="'Material Girl'" />
</template>

<style>
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	/* margin-top: 60px; */
	margin: 0px;
	padding: 0px;

	width: 100vw;
	min-height: 100vh;
	/* background-color: #c3a6ff; */
	background-image: linear-gradient(45deg, #c3a6ff, #f4f0ff);
	box-sizing: border-box;
}
</style>
