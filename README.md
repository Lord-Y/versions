# versions

Versions stand to expose on the UI what we get from `versions-api`.

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
export BASE_URL=http://localhost:8080
export API_URL=http://localhost:8081
npm run ssr:serve
```

### Tailwindcss

If you want to regenerate css file `public/statics/assets/css/css.css`, you can run `npm run tailwind`

### Compiles and minifies for production
```
npm run ssr:build
```

### Start the application
```
npm run ssr:start
```

### Update package

```bash
npm outdated
npm update
```

### Upgrade package to latest

```bash
sudo npm install -g npm-check-updates
ncu
ncu -u
```
