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
        background: var(--body-bg-color);
        color: var(--body-text-color);
    }

    @media (prefers-reduced-motion) {
        :global(html) {
            scroll-behavior: auto;
        }
    }

    #app-root {
        position: relative;
    }
</style>