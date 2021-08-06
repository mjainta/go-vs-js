import { Controller, Get, Param } from '@nestjs/common';
import { DocsService } from './docs.service';
import { Doc } from '../schemas/doc.schema';

@Controller('documents')
export class DocsController {
  constructor(private readonly docsService: DocsService) {}

  @Get()
  getAllDocuments(): Promise<Doc[]> {
    return this.docsService.findAll();
  }

  @Get(':name')
  getFilteredDocuments(@Param('name') name: string): Promise<Doc[]> {
    return this.docsService.findFiltered(name);
  }
}
