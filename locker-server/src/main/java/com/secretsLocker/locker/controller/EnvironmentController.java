package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.copy.CopyEnv;
import com.secretsLocker.locker.dto.delete.DeleteEnvDto;
import com.secretsLocker.locker.dto.path.RepoEnvPath;
import com.secretsLocker.locker.dto.diff.MissingRequest;
import com.secretsLocker.locker.dto.path.RepoPath;
import com.secretsLocker.locker.dto.rename.RenameEnvDto;
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
            @RequestBody RepoEnvPath repoEnvPath
    ) {
        environmentService.create(repoEnvPath);
        return new MessageResponse("EC200", "Environment created.");
    }

    @PostMapping("/get")
    public Response get(@RequestBody RepoPath repoPath) {
        List<String> list = environmentService.get(repoPath);
        return new ListResponse("RP_GT_200", list);
    }

    @PostMapping("/missing")
    public Response missing(@RequestBody MissingRequest missingRequest) {
        List<String> missingSecretsList = environmentService.missing(missingRequest);
        return new ListResponse("EM200", missingSecretsList);
    }

    @PostMapping("/delete")
    public Response delete(@RequestBody DeleteEnvDto deleteEnvDto) {
        environmentService.delete(deleteEnvDto);
        return new MessageResponse("ED200", "Environment deleted.");
    }

    @PostMapping("/copy")
    public Response copy(@RequestBody CopyEnv copyEnv) {
        environmentService.copy(copyEnv);
        return new MessageResponse("EV_CP_200", "Environment " + copyEnv.path() + " copied to " + copyEnv.targetPath() + ".");
    }

    @PostMapping("/rename")
    public Response rename(@RequestBody RenameEnvDto renameEnvDto) {
        environmentService.rename(renameEnvDto);
        return new MessageResponse("EV_RN_200", "Environment " + renameEnvDto.envName + " renamed to " + renameEnvDto.newEnvName + ".");

    }
}
