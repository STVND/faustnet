<template>
    <div id="filebar">
    <div id="filename">

    </div>
    <div id="actions">
        <div id="view">View</div>
        <div id="download">Contents</div>
    </div>
    </div>
</template>

<style scoped>
#filebar {
    display: grid;
    grid-template-areas: 
    "filename actions";
}

#filename {
    display: flex;
    grid-area: filename;
    justify-content: flex-start;
}

#actions {
    display: grid;
    grid-template-areas: 
    "view download";
}
</style>