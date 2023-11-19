<script lang="ts">
    import ChatInput from "$lib/components/chat-input.svelte"
</script>

<div class="flex flex-col justify-between flex-1 bg-gray-750">
    <div class="overflow-y-auto text-sm text-gray-400">
        {#each Array(10).fill(0) as item}
            <div class="flex py-4 mx-6 my-3 border-t border-gray-700">
                <div class="flex-none">
                    <button>
                        <img
                            class="w-10 h-10 rounded-full"
                            src="https://avatars3.githubusercontent.com/u/22605062?s=460&v=4"
                            alt="avatar"
                        />
                    </button>
                </div>
                <div class="ml-5">
                    <div>
                        <button class="text-white hover:underline">drehimself</button>
                        <span class="ml-1 text-xs text-gray-600">07/19/2019</span>
                    </div>
                    <div>
                        <div>yea hahah</div>
                        <div>some other comments</div>
                    </div>
                </div>
            </div>
        {/each}
    </div>
    <ChatInput />
</div>
