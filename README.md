# go-auto-instrumentation

#### 1. Build go-app image
`docker build -t go-app goaap/.`

#### 2. Deploy go-app image with auto instrumentation and Jaeger in your k8s cluster
`kubectl apply -f k8s/.`

#### 3. Verify the pods

```
kubectl get pods 

go-app-instrumented-bbcb75bd-ld75r         2/2     Running   0               7m55s
jaeger-95fc45888-r8f5s                     1/1     Running   5 (2d12h ago)   27d
```

#### 4. Expose the application port and jaeger UI port
```
kubectl port-forward svc/go-app-service 8080:8080
kubectl port-forward svc/jaeger  16686:16686
```

#### 5. Open go-app in web browser - http://localhost:8080

#### 6. Open jaeger UI in web browser - http://localhost:16686

#### 7. Verify service is created and traces are available in jaeger

<img width="1723" alt="Screenshot 2025-03-30 at 21 59 33" src="https://github.com/user-attachments/assets/d03e7157-98d2-4482-bbb4-a93358b58abf" />
