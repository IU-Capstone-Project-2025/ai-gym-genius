export default defineNuxtRouteMiddleware(async (to, from) => {
  // Skip auth check for the auth page itself
  if (to.path === '/auth') {
    return
  }

  // const { isAuthenticated, checkAuth } = useAuth()
  //
  // // Check if user is authenticated
  // const isAuthed = await checkAuth()
  //
  // if (!isAuthed) {
  //   // User is not authenticated, redirect to auth page
  //   return navigateTo('/auth')
  // }
})