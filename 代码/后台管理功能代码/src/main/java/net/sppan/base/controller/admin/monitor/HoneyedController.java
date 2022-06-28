package net.sppan.base.controller.admin.monitor;

import net.sppan.base.controller.BaseController;
import net.sppan.base.entity.AttackedMsg;
import net.sppan.base.service.IAttackedMsgService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@RequestMapping("/admin/machineData")
public class HoneyedController extends BaseController {

    @Autowired
    private IAttackedMsgService iAttackedMsgService;

    @RequestMapping("/list")
    @ResponseBody
    public Page<AttackedMsg> list() {
//		SimpleSpecificationBuilder<Resource> builder = new SimpleSpecificationBuilder<Resource>();
//		String searchText = request.getParameter("searchText");
//		if(StringUtils.isNotBlank(searchText)){
//			builder.add("name", Operator.likeAll.name(), searchText);
//		}
        Page<AttackedMsg> page = iAttackedMsgService.findAll(getPageRequest());
        return page;
    }
}
