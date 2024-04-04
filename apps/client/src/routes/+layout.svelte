<script>
    import "@app/app.css";
    import Navbar from "@components/Navbar/Navbar.svelte";
    import PageTopGradient from "@components/UI/page_top_gradient.svelte";
</script>

<div id="app-root">
    <PageTopGradient />
    <Navbar />
    <slot></slot>
</div>

<style>
    :global(html) {
        background: var(--theme-color-dark);
        color: var(--light-orange-light);
    }

    #app-root {
        position: relative;
    }
</style>