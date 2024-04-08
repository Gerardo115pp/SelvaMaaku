<script>
    import "@app/app.css";
    import Footer from "@components/Footer/footer.svelte";
    import Navbar from "@components/Navbar/Navbar.svelte";
    import PageTopGradient from "@components/Ui_decorations/page_top_gradient.svelte";
</script>

<div id="app-root">
    <PageTopGradient />
    <Navbar />
    <slot></slot>
    <Footer />
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