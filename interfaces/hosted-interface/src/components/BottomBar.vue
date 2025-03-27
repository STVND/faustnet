<template>
    <div>
        <ul>
            <li id="copyleft">&COPY; </li>
            <li id="bottom-info">FaustNET is available under the GPLv3 license. Generally meaning that users are free to copy, modify, and distribute the software at no obligation to pay or attribute as long as they maintain a GPLv3 license on any distribution.
            </li>
        </ul>
    </div>
</template>

<style scoped>
ul {
    display: flex;
    list-style: none;
    justify-content: center;
    height: auto;
}

li {
    display: inline-block;
    padding-left: 3px;
    padding-right: 3px;
    margin-left: 5px;
    margin-right: 5px;
}

#copyleft {
    /* transform: rotate(.5turn); */
    transform: scaleX(-1);
}

</style>