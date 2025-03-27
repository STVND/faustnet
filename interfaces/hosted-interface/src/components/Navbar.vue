<template>
    <div id="navbar">
        <ul>
            <li id="home-link"><RouterLink to="/">Home</RouterLink></li>
            <li id="about-link"><RouterLink to="/about">About</RouterLink></li>
        </ul>
    </div>
</template>


<script>
</script>

<style scoped> 
ul {
    display: flex;
    list-style: none;
    justify-content: space-evenly;
}

li {
    display: inline-block;
    padding-left: 3px;
    padding-right: 3px;
    margin-left: 5px;
    margin-right: 5px;
    outline: double;
    outline-color: #789978a0;
    border-radius: 10px;
}

#navbar {
    font-size: large;
    height: auto;
}

</style>