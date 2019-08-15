import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import About from './views/About.vue';
import Post from './views/PostVideo.vue';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
        },
        {
            path: '/post',
            name: 'psot',
            component: Post,
        },
        {
            path: '/video/:videoID',
            name: 'showVideo',
            component: () => import(/* webpackChunkName: "video" */ './views/ShowVideo.vue'),
        },
        {
            path: '/about',
            name: 'about',
            component: About,
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            // component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
        },
    ],
});