<div class="smk-page-gradient-wrapper">
    <div class="smk-page-gradient"></div>
</div>

<style>
    .smk-page-gradient-wrapper {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height:  var(--navbar-height);
        max-height: var(--navbar-height);
        z-index: var(--z-index-b-2);
    }

    .smk-page-gradient {
        --gradient-color: hsl(from var(--theme-color-light-active) h s l / 0.12 );
        height: 70vh;
        width: 100%;
        background: linear-gradient(180deg, var(--gradient-color) 0%, transparent 100%);
    }

    :global([data-gradient-type="main-page"] .smk-page-gradient) {
        /* --gradient-mix: hsl(162.86,9.59%,28.63%); */
        --gradient-mix: hsl(160, 16%, 25%);
        height: 150vh;
        background: linear-gradient(180deg, var(--gradient-mix) 0%, transparent 80%);
    }
</style>