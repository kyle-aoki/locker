package com.secretsLocker.locker.interceptor;

import com.secretsLocker.locker.service.UserService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@Component
public class AuthInterceptor implements HandlerInterceptor {

    Logger logger = LoggerFactory.getLogger(AuthInterceptor.class);

    @Autowired
    UserService userService;

    @Override
    public boolean preHandle(
            HttpServletRequest request,
            HttpServletResponse response,
            Object handler
    ) throws Exception {
        logger.info("Route '" + request.getRequestURI() + "' is not public. " +
                "Therefore, checking session token before proceeding."
        );

        String username = request.getHeader("username");
        String sessionTokenString = request.getHeader("sessiontoken");

        userService.authenticate(username, sessionTokenString);

        return true;
    }
}
