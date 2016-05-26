package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsElasticTranscoderPreset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsElasticTranscoderPresetCreate,
		Read:   resourceAwsElasticTranscoderPresetRead,
		// Update: resourceAwsElasticTranscoderPresetUpdate,
		Delete: resourceAwsElasticTranscoderPresetDelete,

		Schema: map[string]*schema.Schema{
			"arn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"audio": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					// elastictranscoder.AudioParameters
					Schema: map[string]*schema.Schema{
						"audio_packing_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"bit_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"channels": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"codec_options": &schema.Schema{
							Type:     schema.TypeSet,
							MaxItems: 1,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bit_depth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"bit_order": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"profile": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"signed": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"sample_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"container": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"thumbnails": &schema.Schema{
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					// elastictranscoder.Thumbnails
					Schema: map[string]*schema.Schema{
						"aspect_ratio": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_height": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_width": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"padding_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resolution:": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sizing_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"video": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					// elastictranscoder.VideoParameters
					Schema: map[string]*schema.Schema{
						"aspect_ratio": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"bit_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"codec_options": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
						"display_apect_ratio": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"fixed_gop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"frame_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"key_frames_max_dist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_frame_rate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_height": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"max_width": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"padding_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resolution": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sizing_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"watermarks": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								// elastictranscoder.PresetWatermark
								Schema: map[string]*schema.Schema{
									"horizontal_align": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"horizaontal_offset": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"max_height": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"max_width": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"opacity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"sizing_policy": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"vertical_align": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"vertical_offset": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceAwsElasticTranscoderPresetCreate(d *schema.ResourceData, meta interface{}) error {
	elastictranscoderconn := meta.(*AWSClient).elastictranscoderconn

	req := &elastictranscoder.CreatePresetInput{
		Audio:       expandETAudioParams(d),
		Container:   aws.String(d.Get("container").(string)),
		Description: getStringPtr(d, "description"),
		Name:        aws.String(d.Get("name").(string)),
		Thumbnails:  expandETThumbnails(d),
		Video:       exapndETVideoParams(d),
	}

	log.Printf("[DEBUG] Elastic Transcoder Preset create opts: %s", req)
	resp, err := elastictranscoderconn.CreatePreset(req)
	if err != nil {
		return fmt.Errorf("Error creating Elastic Transcoder Preset: %s", err)
	}

	if resp.Warning != nil && *resp.Warning != "" {
		log.Printf("[WARN] Elastic Transcoder Preset: %s", *resp.Warning)
	}

	d.SetId(*resp.Preset.Id)
	d.Set("arn", *resp.Preset.Arn)

	// there is no Update for presets, so go directly to Read
	return resourceAwsElasticTranscoderPresetRead(d, meta)
}

func expandETThumbnails(d *schema.ResourceData) *elastictranscoder.Thumbnails {
	set, ok := d.GetOk("thumbnails")
	if !ok {
		return nil
	}

	s := set.(*schema.Set)
	if s == nil || s.Len() == 0 {
		return nil
	}
	t := s.List()[0].(map[string]interface{})

	return &elastictranscoder.Thumbnails{
		AspectRatio:   getStringPtr(t, "aspect_ratio"),
		Format:        getStringPtr(t, "format"),
		Interval:      getStringPtr(t, "interval"),
		MaxHeight:     getStringPtr(t, "max_height"),
		MaxWidth:      getStringPtr(t, "max_width"),
		PaddingPolicy: getStringPtr(t, "padding_policy"),
		Resolution:    getStringPtr(t, "resolution"),
		SizingPolicy:  getStringPtr(t, "sizing_policy"),
	}
}

func expandETAudioParams(d *schema.ResourceData) *elastictranscoder.AudioParameters {
	set, ok := d.GetOk("audio")
	if !ok {
		return nil
	}

	s := set.(*schema.Set)
	if s == nil || s.Len() == 0 {
		return nil
	}
	audio := s.List()[0].(map[string]interface{})

	return &elastictranscoder.AudioParameters{
		AudioPackingMode: getStringPtr(audio, "audio_packing_mode"),
		BitRate:          getStringPtr(audio, "bit_rate"),
		Channels:         getStringPtr(audio, "channels"),
		Codec:            getStringPtr(audio, "codec"),
		CodecOptions:     expandETAudioCodecOptions(audio["codec_options"].(*schema.Set)),
		SampleRate:       getStringPtr(audio, "sample_rate"),
	}
}

func expandETAudioCodecOptions(s *schema.Set) *elastictranscoder.AudioCodecOptions {
	if s == nil || s.Len() == 0 {
		return nil
	}

	codec := s.List()[0].(map[string]interface{})

	codecOpts := &elastictranscoder.AudioCodecOptions{
		BitDepth: getStringPtr(codec, "bit_depth"),
		BitOrder: getStringPtr(codec, "bit_prder"),
		Profile:  getStringPtr(codec, "profile"),
		Signed:   getStringPtr(codec, "signed"),
	}

	return codecOpts
}

func exapndETVideoParams(d *schema.ResourceData) *elastictranscoder.VideoParameters {
	set, ok := d.GetOk("video")
	if !ok {
		return nil
	}

	s := set.(*schema.Set)
	if s == nil || s.Len() == 0 {
		return nil
	}
	p := s.List()[0].(map[string]interface{})

	return &elastictranscoder.VideoParameters{
		AspectRatio:        getStringPtr(p, "aspect_ratio"),
		BitRate:            getStringPtr(p, "bit_rate"),
		Codec:              getStringPtr(p, "codec"),
		CodecOptions:       stringMapToPointers(p["codec_options"].(map[string]interface{})),
		DisplayAspectRatio: getStringPtr(p, "display_aspect_ratio"),
		FixedGOP:           getStringPtr(p, "fixed_gop"),
		FrameRate:          getStringPtr(p, "frame_rate"),
		KeyframesMaxDist:   getStringPtr(p, "key_frames_max_dist"),
		MaxFrameRate:       getStringPtr(p, "max_frame_rate"),
		MaxHeight:          getStringPtr(p, "max_height"),
		MaxWidth:           getStringPtr(p, "max_width"),
		PaddingPolicy:      getStringPtr(p, "padding_policy"),
		Resolution:         getStringPtr(p, "resolution"),
		SizingPolicy:       getStringPtr(p, "sizing_policy"),
		Watermarks:         expandETWatermarks(p["watermarks"].(*schema.Set)),
	}
}

func expandETWatermarks(s *schema.Set) []*elastictranscoder.PresetWatermark {
	var watermarks []*elastictranscoder.PresetWatermark

	for _, w := range s.List() {
		watermark := &elastictranscoder.PresetWatermark{
			HorizontalAlign:  getStringPtr(w, "horizontal_align"),
			HorizontalOffset: getStringPtr(w, "horizontal_offset"),
			Id:               getStringPtr(w, "id"),
			MaxHeight:        getStringPtr(w, "max_height"),
			MaxWidth:         getStringPtr(w, "max_width"),
			Opacity:          getStringPtr(w, "opacity"),
			SizingPolicy:     getStringPtr(w, "sizing_policy"),
			Target:           getStringPtr(w, "target"),
			VerticalAlign:    getStringPtr(w, "vertical_align"),
			VerticalOffset:   getStringPtr(w, "vertical_offset"),
		}
		watermarks = append(watermarks, watermark)
	}

	return watermarks
}

func resourceAwsElasticTranscoderPresetUpdate(d *schema.ResourceData, meta interface{}) error {
	panic("not implemented")
}

func resourceAwsElasticTranscoderPresetRead(d *schema.ResourceData, meta interface{}) error {
	elastictranscoderconn := meta.(*AWSClient).elastictranscoderconn

	resp, err := elastictranscoderconn.ReadPreset(&elastictranscoder.ReadPresetInput{
		Id: aws.String(d.Id()),
	})

	if err != nil {
		if err, ok := err.(awserr.Error); ok && err.Code() == "ResourceNotFoundException" {
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[DEBUG] Elastic Transcoder Preset Read response: %#v", resp)

	preset := resp.Preset
	d.Set("arn", *preset.Arn)

	if preset.Audio != nil {
		d.Set("audio", flattenETAudioParameters(preset.Audio))
	}

	d.Set("container", *preset.Container)
	d.Set("name", *preset.Name)

	if preset.Thumbnails != nil {
		d.Set("thumbnails", flattenETThumbnails(preset.Thumbnails))
	}

	d.Set("type", *preset.Type)

	if preset.Video != nil {
		d.Set("video", flattenETVideoParams(preset.Video))
	}

	return nil
}

func flattenETAudioParameters(audio *elastictranscoder.AudioParameters) []map[string]interface{} {
	m := setMap(make(map[string]interface{}))

	m.SetString("audio_packing_mode", audio.AudioPackingMode)
	m.SetString("bit_rate", audio.BitRate)
	m.SetString("channels", audio.Channels)
	m.SetString("codec", audio.Codec)
	m.Set("codec_options", flattenETAudioCodecOptions(audio.CodecOptions))
	m.SetString("sample_rate", audio.SampleRate)

	return m.MapList()
}

func flattenETAudioCodecOptions(opts *elastictranscoder.AudioCodecOptions) []map[string]interface{} {
	if opts == nil {
		return nil
	}

	m := setMap(make(map[string]interface{}))

	m.SetString("bit_depth", opts.BitDepth)
	m.SetString("bit_order", opts.BitOrder)
	m.SetString("profile", opts.Profile)
	m.SetString("signed", opts.Signed)

	return m.MapList()
}

func flattenETThumbnails(thumbs *elastictranscoder.Thumbnails) []map[string]interface{} {
	m := setMap(make(map[string]interface{}))

	m.SetString("aspect_ratio", thumbs.AspectRatio)
	m.SetString("format", thumbs.Format)
	m.SetString("thumbs", thumbs.Interval)
	m.SetString("max_height", thumbs.MaxHeight)
	m.SetString("max_width", thumbs.MaxWidth)
	m.SetString("padding_policy", thumbs.PaddingPolicy)
	m.SetString("resolution", thumbs.Resolution)
	m.SetString("sizing_policy", thumbs.SizingPolicy)

	return m.MapList()
}

func flattenETVideoParams(video *elastictranscoder.VideoParameters) []map[string]interface{} {
	m := setMap(make(map[string]interface{}))

	m.SetString("aspect_ratio", video.AspectRatio)
	m.SetString("bit_rate", video.BitRate)
	m.SetString("codec", video.Codec)
	m.SetStringMap("codec_options", video.CodecOptions)
	m.SetString("display_aspect_ratio", video.DisplayAspectRatio)
	m.SetString("fixed_gop", video.FixedGOP)
	m.SetString("frame_rate", video.FrameRate)
	m.SetString("key_frames_max_dist", video.KeyframesMaxDist)
	m.SetString("max_frame_rate", video.MaxFrameRate)
	m.SetString("max_height", video.MaxHeight)
	m.SetString("max_width", video.MaxWidth)
	m.SetString("padding_policy", video.PaddingPolicy)
	m.SetString("resolution", video.Resolution)
	m.Set("watermarks", flattenETWatermarks(video.Watermarks))

	return m.MapList()
}

func flattenETWatermarks(watermarks []*elastictranscoder.PresetWatermark) []map[string]interface{} {
	var watermarkSet []map[string]interface{}

	for _, w := range watermarks {
		watermark := setMap(make(map[string]interface{}))

		watermark.SetString("horizontal_align", w.HorizontalAlign)
		watermark.SetString("horizontal_offset", w.HorizontalOffset)
		watermark.SetString("id", w.Id)
		watermark.SetString("max_height", w.MaxHeight)
		watermark.SetString("max_width", w.MaxWidth)
		watermark.SetString("opacity", w.Opacity)
		watermark.SetString("sizing_policy", w.SizingPolicy)
		watermark.SetString("target", w.Target)
		watermark.SetString("vertical_align", w.VerticalAlign)
		watermark.SetString("vertical_offset", w.VerticalOffset)

		watermarkSet = append(watermarkSet, watermark.Map())
	}

	return watermarkSet
}

func resourceAwsElasticTranscoderPresetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
