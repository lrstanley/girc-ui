<script lang="ts">
    import "./styles.css"
    import ServerSidebar from "$lib/components/server-sidebar.svelte"
    import WindowTopbar from "$lib/components/window-topbar.svelte"
    import ChannelSidebar from "$lib/components/channel-sidebar.svelte"
    import ChatHistory from "$lib/components/chat-history.svelte"
    import UserSidebar from "$lib/components/user-sidebar.svelte"
</script>

<div class="flex">
    <ServerSidebar />
    <div class="flex flex-col flex-1 h-screen min-h-screen">
        <WindowTopbar />

        <div class="flex flex-1 overflow-y-hidden bg-blue-500">
            <ChannelSidebar />
            <div class="flex justify-between flex-1">
                <ChatHistory />
                <UserSidebar />
            </div>
        </div>
    </div>
</div>
