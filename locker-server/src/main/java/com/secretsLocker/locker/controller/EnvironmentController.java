package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateEnvironmentDto;
import com.secretsLocker.locker.dto.GetEnvDto;
import com.secretsLocker.locker.dto.KeyValue;
import com.secretsLocker.locker.dto.diff.MissingRequest;
import com.secretsLocker.locker.response.KeyValueResponse;
import com.secretsLocker.locker.response.ListResponse;
import com.secretsLocker.locker.response.MessageResponse;
import com.secretsLocker.locker.response.Response;
import com.secretsLocker.locker.service.EnvironmentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/env")
public class EnvironmentController {

    @Autowired
    EnvironmentService environmentService;

    @PostMapping("/create")
    public Response createEnv(
            @RequestBody CreateEnvironmentDto createEnvironmentDto
    ) {
        environmentService.create(createEnvironmentDto);
        return new MessageResponse("EC200", "Environment created.");
    }

    @PostMapping("/get")
    public Response get(@RequestBody GetEnvDto getEnvDto) {
        List<KeyValue> keyValues = environmentService.get(getEnvDto);
        return new KeyValueResponse("EG200", keyValues);
    }

    @PostMapping("/missing")
    public Response missing(@RequestBody MissingRequest missingRequest) {
        List<String> missingSecretsList = environmentService.missing(missingRequest);
        return new ListResponse("EM200", missingSecretsList);
    }
}
