import { createWebHistory, createRouter } from "vue-router"
import Home from "../components/Home.vue"
import SignUp from "@/components/SignUp.vue";
import SignIn from "@/components/SignIn.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path: "/signup",
        name: "SignUp",
        component: SignUp,
    },
    {
        path: "/signin",
        name: "SignIn",
        component: SignIn,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router