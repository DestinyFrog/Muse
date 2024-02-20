<template>
	<header class="header">
		<h2>Muse</h2>

<!--
		<ul class="menu">
			<li>Songs</li>
			<li>Autors</li>
		</ul>
-->
	</header>
</template>

<script>
export default {
	// data() {

	// }
}
</script>

<style>
.header {
	background-color: #5503af;
	color: #c1e1fd;

	/* width: fit-; */

	margin: 0px;

	display: flex;
	flex-direction: row;

	padding: 0px 10px;

	justify-content: space-between;
}

.header h2 {
	margin: 6px 0px;
}

.menu {
	padding: 0;
	margin: 0;
	list-style: none;
	display: flex;
	height: auto;
}

.menu li {
	width: 80px;
	height: 100%;

	display: flex;
	justify-content: center;
	align-items: center;	

	border-bottom: 1px solid aliceblue;
}

.menu li:hover {
	background-color: #810fff;
}
</style>