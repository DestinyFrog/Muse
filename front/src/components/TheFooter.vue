<template>
	<footer>
		<h1 v-show="SelectedSong">{{ SelectedSong }}</h1>

		<!-- <audio v-show="file" controls> -->
			<!-- <source :src="file" type="audio/mpeg" /> -->
		<!-- </audio> -->
	</footer>

</template>

<script>
export default {
	props: [ 'SelectedSong' ],

	data() {
		return {
			file: null
		}
	}
}
</script>

<style>
	footer {
		position: fixed;
		bottom: 0px;

		width: 100vw;

		background-color: #5503af;
		color: #c1e1fd;
	}
</style>