	testRepo, _, cleanupFn := testhelper.NewTestRepo(t)
			remove, err := operations.OverrideHooks(hookName, hookContent)
			require.NoError(t, err)
			defer remove()