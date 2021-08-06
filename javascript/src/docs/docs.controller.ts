import { Controller, Get } from '@nestjs/common';
import { DocsService } from './docs.service';
import { Doc } from '../schemas/doc.schema';

@Controller('documents')
export class DocsController {
  constructor(private readonly docsService: DocsService) {}

  @Get()
  getHello(): Promise<Doc[]> {
    return this.docsService.findAll();
  }
}
