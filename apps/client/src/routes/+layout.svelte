<script>
    import "@app/app.css";
    import Navbar from "@components/Navbar/Navbar.svelte";
</script>

<Navbar />
<slot></slot>

<style>
    :global(html) {
        background: var(--theme-color-dark);
        color: var(--light-orange-light);
    }
</style>